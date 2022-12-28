FROM cassandra:latest

# Add your custom Cassandra configuration files

# Run the schema.cql script to set up the Cassandra keyspace and tables
CMD ["cassandra", "-f"]


# FROM ubuntu:18.04

# # Install Java and other dependencies
# RUN apt-get update && apt-get install -y openjdk-8-jdk wget tar

# # Download and extract Cassandra
# RUN wget http://archive.apache.org/dist/cassandra/4.1.1/apache-cassandra-4.1.1-bin.tar.gz
# RUN tar -xvzf apache-cassandra-4.1.1-bin.tar.gz

# # Set environment variables and start Cassandra
# ENV CASSANDRA_HOME /apache-cassandra-4.1.1
# ENV PATH $CASSANDRA_HOME/bin:$PATH
# CMD ["cassandra", "-R", "-f"]



# FROM alpine:3.12

# RUN apk update && apk add wget openjdk8-jre-base

# RUN wget -q -O - https://www.apache.org/dist/cassandra/KEYS | apk add --import -

# RUN echo "http://www.apache.org/dist/cassandra/alpine/3.12/apache-cassandra-3.12-bin.tar.gz" >> /etc/apk/sources.list

# RUN apk update && apk add apache-cassandra

# CMD ["cassandra", "-f"] 