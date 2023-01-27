FROM golang:1.19.5-buster

RUN apt-get update && apt-get install -y openssh-server


RUN mkdir -p /usr/src/launchcode && mkdir /var/run/sshd
RUN echo 'root:root' | chpasswd

COPY docker/sshd_config /etc/ssh/sshd_config
COPY ./ /usr/src/launchcode
COPY docker/copy-src.sh /usr/local/bin/copy-src
COPY docker/build-launchers.sh /usr/local/bin/build-launchers

EXPOSE 22
CMD ["/usr/sbin/sshd", "-D"]
