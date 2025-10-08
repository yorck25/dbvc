#include <iostream>

enum LogType {
    WARNING = 0,
    INFO = 1,
    ERROR = 2
};

const char* LogTypes[] = {"Warning", "Info", "Error"};

class Logger {
public:
    void print(LogType type = LogType::INFO, const std::string& msg = "") {
        std::cout << "[" + std::string(LogTypes[type]) + "]: " + msg << std::endl;
    }

    void print(const std::string& msg) {
        print(LogType::INFO, msg);
    }
};
