FROM alpine:3.7

RUN apk --update add squid

RUN sed -ie s/"3128/3129"/g /etc/squid/squid.conf

# This step enables access from private IP address ranges.
# RUN sed -ie s/"http_access deny all"/"http_access allow all"/g /etc/squid/squid.conf
RUN sed -ie s/"http_access deny !Safe_ports"/"http_access allow !Safe_ports"/g /etc/squid/squid.conf
RUN sed -ie s/"http_access deny CONNECT !SSL_ports"/"http_access allow CONNECT !SSL_ports"/g /etc/squid/squid.conf

RUN echo "acl MINIKUBE dst 192.168.0.0/24" | tee -a /etc/squid/squid.conf
RUN echo "http_access allow MINIKUBE" | tee -a /etc/squid/squid.conf
RUN echo "http_access deny all" | tee -a /etc/squid/squid.conf

CMD ["squid", "-NYCd", "1"]
