FROM grafana/alloy
COPY ./config.alloy /etc/alloy/config.alloy
CMD ["run", "--server.http.listen-addr=0.0.0.0:12345", "--server.http.ui-path-prefix=/alloy", "--storage.path=/var/lib/alloy/data", "--disable-reporting=true", "/etc/alloy/config.alloy"]
