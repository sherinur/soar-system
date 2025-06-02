FROM golang:1.24

WORKDIR /app

COPY . .

EXPOSE 50052 2006

CMD ["./auth-bin"]

# # Uncomment for a diagnostic debug (container-config-error.md shared runbook)
# CMD ["sleep", "infinity"]