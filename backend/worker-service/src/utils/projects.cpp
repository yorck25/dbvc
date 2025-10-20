#include <iostream>
#include "../connectors/PostgresConnector.h"
#include "../logger/Logger.h"

using namespace std;

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

    Project(int id, int owner_id, std::string name, std::string description, std::string created_at,
            std::string updated_at, bool active, std::string visibility, int connection_type)
            : id(id), owner_id(owner_id), name(name), description(description), created_at(created_at),
              updated_at(updated_at), active(active), visibility(visibility), connection_type(connection_type) {}
};

std::vector<Project> getProjects(PostgresConnector *psqlConnector, Logger &logger) {
    std::string query = "SELECT * FROM projects WHERE active = true";

    if (!psqlConnector) {
        logger.print(LogType::ERROR, "Database connector is null");
        return {};
    }

    auto result = psqlConnector->executeQuery(query);

    if (!result) {
        logger.print(LogType::ERROR, "Fail to load projects");
        return {};
    }

    return Mapper::mapProjects(*result);
}