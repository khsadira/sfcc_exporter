########
#README#
########

SFCC(b2c commerce) to prometheus exporter

Installation: (without docker)
1: Make
2: ./sfcc_prom_exporter

Installation: (with docker)
1: Make
1: Make docker
3: docker run -d -p 9240:940 sfcc_exporter

Then you can use it easily by using your navigator at  localhost:9240/metrics/site={site_id} or `ip`:9240/metrics/site={site_id} if you want to ask from another computer

Khan S.
