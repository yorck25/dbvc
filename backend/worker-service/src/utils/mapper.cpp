#include <pqxx/pqxx>
#include "projects.cpp"

class Mapper {
public:
    static Project mapProject(const pqxx::row &row) {
        int id = row["id"].as<int>();
        int owner_id = row["owner_id"].as<int>();
        std::string name = row["name"].as<std::string>();
        std::string description = row["description"].as<std::string>();
        std::string created_at = row["created_at"].as<std::string>();
        std::string updated_at = row["updated_at"].as<std::string>();
        bool active = row["active"].as<bool>();
        std::string visibility = row["visibility"].as<std::string>();
        int connection_type = row["connection_type"].as<int>();

        return Project(
                id,
                owner_id,
                name,
                description,
                created_at,
                updated_at,
                active,
                visibility,
                connection_type
        );
    }

    static std::vector<Project> mapProjects(const pqxx::result &res) {
        std::vector<Project> projects;
        projects.reserve(res.size());
        for (const auto &row: res) {
            projects.push_back(mapProject(row));
        }
        return projects;
    }
};