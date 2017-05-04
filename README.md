# Data Sipper Data Source (Proof of Concept)
A simple proof of concept application connect to distributed data sources across multiple organizations, and send to a centralized database in a public cloud. This project provides an example ETL job to extract data from a local datasource and send to the central server.

# Talend ETL
This example uses the open source version of [Talend](https://www.talend.com/download/talend-open-studio/) to create a JAR file that can be scheduled to run within the local environment on a periodic basis to extract recent data from a datasource. 