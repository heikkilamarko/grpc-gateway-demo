FROM prom/prometheus
COPY prometheus.yml /etc/prometheus/prometheus.yml
CMD ["--config.file=/etc/prometheus/prometheus.yml", "--storage.tsdb.path=/prometheus", "--web.enable-remote-write-receiver", "--enable-feature=exemplar-storage", "--enable-feature=native-histograms"]
