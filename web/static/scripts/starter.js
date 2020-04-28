/**
 * Copyright 2020 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

const PATHWAY_STARTER_CONTAINER_ID = 'pathway-starter';
const REQUEST_TEXTAREA_CONTAINER_CLASS = 'request';
const RESPONSE_TEXT_CONTAINER_CLASS = 'response';

/**
 * Appends the request textarea element to the pathway starter container.
 * @param {!Element} pathwayStarterContainer
 */
function appendRequestTextArea(pathwayStarterContainer) {
  pathwayStarterContainer.append('div').attr(
      HTML_ATTRIBUTES.class, REQUEST_TEXTAREA_CONTAINER_CLASS);
  const requestTextAreaContainer =
      d3.select(`.${REQUEST_TEXTAREA_CONTAINER_CLASS}`);
  requestTextAreaContainer.append('textarea')
      .attr(HTML_ATTRIBUTES.class, 'request-text')
      .attr(
          HTML_ATTRIBUTES.placeholder, 'Enter pathway name or YML definition');
  requestTextAreaContainer.append('input')
      .attr(HTML_ATTRIBUTES.type, 'submit')
      .attr(HTML_ATTRIBUTES.value, 'Run or Send')
      .attr(HTML_ATTRIBUTES.class, 'request-button');
}

/**
 * Appends the response element to the pathway starter container for displaying
 * the response from the backend.
 * @param {!Element} pathwayStarterContainer
 */
function appendResponseTextElement(pathwayStarterContainer) {
  pathwayStarterContainer.append('div').attr(
      HTML_ATTRIBUTES.class, RESPONSE_TEXT_CONTAINER_CLASS);
  const responseTextContainer = d3.select(`.${RESPONSE_TEXT_CONTAINER_CLASS}`);
  responseTextContainer.append('text').attr(
      HTML_ATTRIBUTES.class, 'response-text');
}

/**
 * Sends the specified pathway to the backend to be executed.
 * @param {string} request
 */
function sendRequest(request) {
  d3.request(document.getElementById(PATHWAY_STARTER_CONTAINER_ID).dataset.path)
      .header('X-Requested-With', 'XMLHttpRequest')
      .header('Content-Type', 'application/x-www-form-urlencoded')
      .post(request, function(data) {});
  clearResponse();
  requestResponse();
}

/** Requests response from backend and parses it into a displayable string. */
function requestResponse() {
  d3.request(document.getElementById(PATHWAY_STARTER_CONTAINER_ID).dataset.path)
      .get(function(data) {
        const responseData = data.response;
        if (responseData === '') {
          requestResponse();
        } else {
          const splitResponses = responseData.split('\n');
          const firstLine = splitResponses[0];
          splitResponses.splice(0, 1);
          fillResponse(firstLine, splitResponses);
        }
      });
}

/**
 * Creates an html response text where the first line is standalone and the
 * rest of the supplied lines, if any, are formatted as an html list.
 * @param {string} firstLine The first line in the response
 * @param {?Array<string>} rest The rest of the lines in the response
 */
function fillResponse(firstLine, rest) {
  clearResponse();
  const responseTextContainer = d3.select(`.${RESPONSE_TEXT_CONTAINER_CLASS}`);
  responseTextContainer.select('.response-text').text(firstLine);
  const ul = responseTextContainer.append('text')
                 .attr(HTML_ATTRIBUTES.class, 'textfield-response-name')
                 .append('ul');
  ul.selectAll('li').data(rest).enter().append('li').html(String);
}

/** Clears the response text in the UI. */
function clearResponse() {
  const responseTextContainer = d3.select(`.${RESPONSE_TEXT_CONTAINER_CLASS}`);
  responseTextContainer.selectAll('.textfield-response-name').remove();
  responseTextContainer.select('.response-text').text('');
}

// Select the main pathway starter container and attach the request and response
// elements.
const pathwayStarterContainer = d3.select(`#${PATHWAY_STARTER_CONTAINER_ID}`);
appendRequestTextArea(pathwayStarterContainer);
appendResponseTextElement(pathwayStarterContainer);

// Attach a click handler to the submit button.
d3.select('.request-button').on('click', function() {
  sendRequest(document.getElementsByClassName('request-text')[0].value);
});
