#include <iostream>
#include "connectors/BaseConnector.cpp"
#include "logger/Logger.cpp"
#include <pqxx/pqxx>

class PostgresConnector : public BaseConnector {
public:
    bool tryConnect() override {
        try {
            pqxx::connection c{"dbname=postgres user=postgres password=test1234 host=localhost"};
            pqxx::work txn{c};

            pqxx::result r = txn.exec("SELECT version()");
            std::cout << "PostgreSQL version: " << r[0][0].c_str() << std::endl;

            txn.commit();
        } catch (const std::exception &e) {
            std::cerr << "Error: " << e.what() << std::endl;
            return 1;
        }
        return true;
    }
};

class MySQLConnector : public BaseConnector {
public:
    bool tryConnect() override {
        return true;
    }
};

std::unique_ptr<BaseConnector> getConnector(const std::string &connectionType) {
    if (connectionType == "psql") {
        return std::make_unique<PostgresConnector>();
    } else if (connectionType == "mysql") {
        MySQLConnector mysqlConnector;
        return std::make_unique<MySQLConnector>();
    }

    return nullptr;
}

int main() {
    Logger logger;
    logger.print("Start worker.");

    std::unique_ptr<BaseConnector> connector = getConnector("psql");

    bool res = connector->connect();
    if (!res) {
        logger.print(LogType::ERROR, "Fail to connect to db");
        return 1;
    }
    logger.print("Connected to database");

    return 0;
};