// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package files supports reading and writing files from local directories or GCS.
package files

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

const gcsBucketPrefix = "gs://"

// File represents a file, either local or remote.
type File interface {
	Read(ctx context.Context) ([]byte, error)
	Name() string
	FullPath() string
}

// List lists files in the directory specified by the path.
func List(ctx context.Context, path string) ([]File, error) {
	if strings.HasPrefix(path, gcsBucketPrefix) {
		return listGCSFiles(ctx, path)
	}
	return listLocalFiles(path)
}

// Read reads the file specified by the path.
func Read(ctx context.Context, path string) ([]byte, error) {
	if strings.HasPrefix(path, gcsBucketPrefix) {
		return readGCSFile(ctx, path)
	}
	return readLocalFile(path)
}

func readGCSFile(ctx context.Context, path string) ([]byte, error) {
	f, err := listGCSFiles(ctx, path)
	if err != nil {
		return nil, err
	}
	if len(f) != 1 {
		return nil, fmt.Errorf("%s does not identify a file", path)
	}
	return f[0].Read(ctx)
}

func listGCSFiles(ctx context.Context, path string) ([]File, error) {
	b, prefix, err := parseGCSPath(path)
	if err != nil {
		return nil, err
	}
	c, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	bucket := c.Bucket(b)
	it := bucket.Objects(ctx, &storage.Query{Prefix: prefix})
	var files []File
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		files = append(files, gcsFile{prefix: prefix, object: bucket.Object(attrs.Name)})
	}
	return files, nil
}

// parseGCSPath accepts a path of the form gs://bucket/path/to/dir and returns
// ("bucket", "path/to/dir", <error>). A path with just the bucket name is also
// valid (i.e. gs://bucket)
func parseGCSPath(path string) (string, string, error) {
	if !strings.HasPrefix(path, gcsBucketPrefix) {
		return "", "", fmt.Errorf("GCS path has an invalid format: %s", path)
	}
	p := strings.TrimPrefix(path, gcsBucketPrefix)
	i := strings.Index(p, "/")
	if i == -1 {
		return p, "", nil
	}
	return p[:i], p[i+1:], nil
}

type gcsFile struct {
	prefix string
	object *storage.ObjectHandle
}

func (f gcsFile) Name() string {
	return strings.TrimPrefix(f.object.ObjectName(), fmt.Sprintf("%s/", f.prefix))
}

func (f gcsFile) FullPath() string {
	return f.object.ObjectName()
}

func (f gcsFile) Read(ctx context.Context) ([]byte, error) {
	r, err := f.object.NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return ioutil.ReadAll(r)
}

func listLocalFiles(path string) ([]File, error) {
	dirFiles, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var files []File
	for _, f := range dirFiles {
		if f.IsDir() {
			continue
		}
		files = append(files, localFile{path, f.Name()})
	}
	return files, nil
}

type localFile struct {
	dirName  string
	fileName string
}

func (f localFile) Name() string {
	return f.fileName
}

func (f localFile) FullPath() string {
	return path.Join(f.dirName, f.fileName)
}

func (f localFile) Read(_ context.Context) ([]byte, error) {
	return readLocalFile(f.FullPath())
}

func readLocalFile(path string) ([]byte, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if fi.IsDir() {
		return nil, fmt.Errorf("can't read local file; %s is a directory", path)
	}
	return ioutil.ReadFile(path)
}
