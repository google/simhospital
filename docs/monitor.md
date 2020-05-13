# Monitor Simulated Hospital

-   [Multiple instances](#multiple-instances)

You can use [Prometheus](https://prometheus.io/) to browse and monitor metrics
from Simulated Hospital. Prometheus sends HTTP requests to pull data from
Simulated Hospital.

Simulated Hospital starts an HTTP server on port 9095 which serves
Prometheus-compatible metric data on the path: `/metrics`. To see which signals
you can monitor, open the following page in your browser (after changing the
hostname to match your computer): `http://host.example.com:9095/metrics`.

Add your server with the `/metrics` path to your Prometheus targets. To learn
more, visit the [Prometheus documentation site](https://prometheus.io/docs/).

## Multiple instances

If you're running more than one instance of Simulated Hospital, you can change
the TCP port that serves metrics. Set the port with the `metrics_listen_address`
command-line argument when you launch Simulated Hospital. To learn more, visit
[Command-line arguments](./arguments.md#runtime).
