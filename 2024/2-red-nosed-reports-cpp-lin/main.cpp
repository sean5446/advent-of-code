/*
 Compile from command line:
 g++ -o main.exe main.cpp
*/

#include <iostream>
#include <fstream>
#include <sstream>
#include <vector>
#include <string>
#include <algorithm>

using namespace std;

// TODO optimzaition - for part1, we don't need a vector of strings, we can just read the file line by line


vector<string> readLines(const string& fileLocation) {
    vector<string> lines;
    ifstream file(fileLocation);

    if (!file.is_open()) {
        cerr << "Could not open the file: " << fileLocation << endl;
        return lines;
    }

    string line;
    while (getline(file, line)) {
        lines.push_back(line);
    }

    file.close();
    return lines;
}

int processFileAndCompare1(const string& fileLocation) {
    vector<string> lines = readLines(fileLocation);
    
    int sum = 0;

    for (const auto& line : lines) {
        istringstream stream(line);
        int previousValue = -1;
        int currentValue;

        int increasing = 0;
        int decreasing = 0;
        bool isInvalid = false;

        while (stream >> currentValue) {
            // if on first element, continue to next
            if (previousValue == -1) {
                previousValue = currentValue;
                continue;
            }
            // compare current and previous values
            if (currentValue == previousValue) {
                isInvalid = true;
                break;
            } else if (currentValue > previousValue) {
                if (currentValue - previousValue > 3) {
                    isInvalid = true;
                    break;
                }
                increasing += 1;
            } else { // currentValue < previousValue
                if (previousValue - currentValue > 3) {
                    isInvalid = true;
                    break;
                }
                decreasing += 1;
            }
            previousValue = currentValue;
        }
        if (!isInvalid && ((increasing > 0 && decreasing == 0) || (increasing == 0 && decreasing > 0))) {
            cout << "Valid: " << line << "\n";
            sum += 1;
        }
    }

    return sum;
}

bool isGood(vector<int>& values) {
    // check that no two values have a difference greater than 3
    for (int i = 0; i < values.size() - 1; i++) {
        int diff = abs(values[i] - values[i + 1]);
        if (diff > 3 || diff < 1) {
            cout << values[i] << " " << values[i + 1] << " diff " << diff << endl;
            return false;
        }
    }

    // make sure they are increasing or decreasing
    vector<int> sorted = values;
    vector<int> revsorted = values;
    sort(sorted.begin(), sorted.end());
    sort(revsorted.begin(), revsorted.end(), greater<int>());
    if (values == sorted || values == revsorted) {
        cout << "not sorted!" << endl;
        return true;
    }

    return false;
}


int processFileAndCompare2(const string& fileLocation) {
    vector<string> lines = readLines(fileLocation);
    
    int sum = 0;

    for (const auto& line : lines) {
        istringstream stream(line);
        int currentValue;
        vector <int> values;
        while (stream >> currentValue) {
            values.push_back(currentValue);
        }

        cout << "Checking: " << line << "\n";

        if (isGood(values)) {
            cout << "Valid: " << line << "\n";
            sum += 1;
        } else {
            // remove each element and check if it is still good
            for (int i = 0; i < values.size(); i++) {
                vector<int> copy = values;
                copy.erase(copy.begin() + i);
                if (isGood(copy)) {
                    cout << "Valid: " << line << " (removed " << values[i] << ")\n";
                    sum += 1;
                    break;
                }
            }
        }
        // string inputLine;
        // getline(cin, inputLine);
    }
    return sum;
}

int main() {
    int sum = processFileAndCompare2("input.txt");
    cout << "Sum: " << sum << endl;

    return 0;
}

