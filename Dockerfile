FROM plugins/base:multiarch

LABEL maintainer="Leyou Hong <lawrencehongengineer@gmail.com>" \
  org.label-schema.name="Drone Workshop" \
  org.label-schema.vendor="Leyou Hong" \
  org.label-schema.schema-version="1.0"

ADD release/linux/amd64/telo /bin/

HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD [ "/bin/telo", "-ping" ]

ENTRYPOINT ["/bin/telo"]
