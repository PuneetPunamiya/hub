FROM registry.access.redhat.com/ubi8/nodejs-16:1-8 as BUILD

WORKDIR /opt/app-root/src

USER root
# install deps and the copy the src to speed up docker build
COPY package-lock.json package.json tsconfig.json /opt/app-root/src/
RUN npm install

COPY public /opt/app-root/src/public/
COPY src /opt/app-root/src/src/
RUN npm run build

# Stage 2 - the production environment
FROM registry.access.redhat.com/ubi8/nginx-120:1-7
COPY --from=BUILD /opt/app-root/src/build /usr/share/nginx/html
COPY image/start.sh /usr/bin/

USER root
RUN chmod ugo+rw /usr/share/nginx/html/config.js  && \
    chown nginx:nginx /usr/share/nginx/html/config.js && \
    chmod +x /usr/bin/start.sh
USER nginx

EXPOSE 8080

COPY image/nginx.conf "${NGINX_CONFIGURATION_PATH}"

CMD /usr/bin/start.sh
