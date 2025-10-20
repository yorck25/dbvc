#pragma once
#include <pqxx/pqxx>
#include <optional>
#include <memory>
#include <string>

class PostgresConnector {
public:
    PostgresConnector(std::string host, int port, std::string user,
                      std::string password, std::string database);

    bool connect();
    std::optional<pqxx::result> executeQuery(const std::string &query);

private:
    bool tryConnect();

    std::string host;
    int port;
    std::string user;
    std::string password;
    std::string database;
    int maxRetries = 5;

    std::unique_ptr<pqxx::connection> connection;
};
