FROM woodenfish42/hustoj-base as build-env
LABEL maintainer="muzea <mr.muzea@gmail.com>"

RUN set -ex \
    && apt install \
               gcc \
               g++ \
               make \
               libmysqlclient-dev \
               libmysql++-dev -y \
    && mkdir /home/judge \
    && cd /hustoj/trunk/core/judged \
    && make \
    && cd ../judge_client \
    && make

FROM debian:8

RUN set -ex \
    && apt update \
    && apt install \
               gcc \
               g++ \
               libmysqlclient-dev \
               libmysql++-dev -y \
    && mkdir /home/judge \
    && mkdir /home/judge/etc \
    && mkdir /home/judge/data \
    && mkdir /home/judge/log \
    && mkdir /home/judge/run0 \
    && useradd -m -u 1536 judge \
    && chown -R judge /home/judge \
    && chmod 775 /home/judge /home/judge/data /home/judge/etc /home/judge/run?

COPY --from=build-env /hustoj/trunk/core/judged/judged /usr/bin/judged
COPY --from=build-env /hustoj/trunk/core/judge_client/judge_client /usr/bin/judge_client

RUN set -ex \
    && chmod +x /usr/bin/judged \
    && chmod +x /usr/bin/judge_client

CMD judged && tail -f /dev/null