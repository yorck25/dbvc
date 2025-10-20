#pragma once
#include <string>
#include "../connectors/PostgresConnector.h"
#include "../logger/Logger.h"

class Project {
public:
    int id;
    int owner_id;
    std::string name;
    std::string description;
    std::string created_at;
    std::string updated_at;
    bool active;
    std::string visibility;
    int connection_type;

    Project(int id,
            int owner_id,
            std::string name,
            std::string description,
            std::string created_at,
            std::string updated_at,
            bool active,
            std::string visibility,
            int connection_type);
};

std::vector<Project> getProjects(PostgresConnector *psqlConnector, Logger &logger)