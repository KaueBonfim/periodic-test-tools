from ubuntu

EXPOSE 5080
COPY ./main .
RUN chmod +X /main
CMD ["./main"]
