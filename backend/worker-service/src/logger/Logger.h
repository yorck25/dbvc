#pragma once
#include <iostream>
#include <string>

enum class LogType {
    WARNING = 0,
    INFO = 1,
    ERROR = 2
};

class Logger {
public:
    void print(LogType type, const std::string& msg);
    void print(const std::string& msg);
};
