FROM alpine
ADD smallMicro /smallMicro
ENTRYPOINT [ "/smallMicro" ]
