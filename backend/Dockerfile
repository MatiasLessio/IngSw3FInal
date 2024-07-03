FROM golang:1.18

EXPOSE 8080

# Set working directory
WORKDIR /backend

# Add source code
ADD . /backend

# Install dependencies
RUN go mod tidy

# Install Cloud SQL Proxy
RUN wget https://dl.google.com/cloudsql/cloud_sql_proxy.linux.amd64 -O cloud_sql_proxy && chmod +x cloud_sql_proxy

# Build the application
RUN go build -o backend

# Make sure the application and proxy are executable
RUN chmod +x /backend && chmod +x cloud_sql_proxy

# Start Cloud SQL Proxy and the application
CMD ./cloud_sql_proxy -instances=ingsw3-428221:us-central1:ingsw3=tcp:3306 & ./backend
