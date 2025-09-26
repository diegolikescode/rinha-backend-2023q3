# rinha-backend-2023q3

## requirements (just for metrics)

- node_exporter from Prometheus on your binaries
    - create a new user: ``useradd  --no-create-home --shell /bin/false node_exporter``
    - give the bin to the user: ``chown node_exporter:node_exporter /usr/local/bin/node_exporter`` (run ``which node_exporter`` if necessary)
    - copy node_exporter.service: cp node_exporter.service ``/etc/systemd/system/node_exporter.service``
    - reload the daemon: ``systemctl daemon-reload``
    - enable service: ``systemctl enable node_exporter``
    - start: ``systemctl start node_exporter``
    - check if all went well: ``systemctl status node_exporter.service``

## some numbers

- simpler-version: ~41,500
