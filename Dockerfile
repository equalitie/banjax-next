# Copyright (c) 2020, eQualit.ie inc.
# All rights reserved.
# 
# This source code is licensed under the BSD-style license found in the
# LICENSE file in the root directory of this source tree.

FROM golang:1.15.0-buster

RUN set -x \
 && DEBIAN_FRONTEND=noninteractive apt-get update \
 && DEBIAN_FRONTEND=noninteractive apt-get install -y \
		iptables

RUN mkdir -p /opt/banjax-next
COPY ./ /opt/banjax-next/
RUN cd /opt/banjax-next && go test

RUN mkdir -p /etc/banjax-next
COPY ./banjax-next-config.yaml /etc/banjax-next/
# COPY ./caroot.pem /etc/banjax-next/
# COPY ./certificate.pem /etc/banjax-next/
# COPY ./key.pem /etc/banjax-next/

RUN mkdir -p /var/log/banjax-next

EXPOSE 8081

WORKDIR /opt/banjax-next

CMD ["go", "run", "banjax-next.go"]
