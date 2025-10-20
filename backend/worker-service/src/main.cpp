#include <iostream>
#include "connectors/PostgresConnector.h"
#include "connectors/BaseConnector.cpp"
#include "logger/Logger.h"
#include <pqxx/pqxx>

using namespace std;

void processor(const Project& project){
    cout << "Project = " << project.name << endl;

}

int main() {
    Logger logger;
    logger.print("Start worker.");

    PostgresConnector psqlConnector = PostgresConnector(
            "localhost", 5432, "postgres", "test1234", "postgres"
    );

    if (!psqlConnector.connect()) {
        logger.print(LogType::ERROR, "Fail to connect to db");
        return 1;
    }
    logger.print("Connected to database");

    std::vector<Project> projects = getProjects(&psqlConnector, logger);

    for (const Project& project : projects) {
        processor(project);
    }

    return 0;
}