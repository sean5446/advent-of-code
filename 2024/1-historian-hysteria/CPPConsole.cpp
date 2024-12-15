
/*
 Compilation from command line:
 cl CPPConsole.cpp
*/

#include <iostream> // cout, cerr
#include <fstream> // file stream
#include <sstream> // string stream
#include <algorithm> // sort

#include <string>
#include <vector>

#include <windows.h>  // GetModuleFileNameA

using namespace std;

string getCurrentPathWindows() {
    char path[MAX_PATH];
    DWORD result = GetModuleFileNameA(NULL, path, MAX_PATH);
    if (result != 0) {
        string a = path;
        size_t pos = a.find_last_of("\\");
        if (pos != string::npos) {
            return a.substr(0, pos);
        }
    }
	return "";
}

void readFileLines(const string& filePath, vector<int>& left, vector<int>& right) {
    ifstream file(filePath);

    if (!file.is_open()) {
        cerr << "Error opening file!" << endl;
        return;
    }

    string line;

    while (getline(file, line)) {
        istringstream stream(line);
        string word;
        if (stream >> word) {
            left.push_back(stoi(word));
        }
        if (stream >> word) {
            right.push_back(stoi(word));
        }
    }

    file.close();
}


int problem1() {
    string fileName = getCurrentPathWindows() + "\\input.txt";
	cout << "Reading file: " << fileName << "\n";

    vector<int> left;
    vector<int> right;
    readFileLines(fileName, left, right);

    sort(left.begin(), left.end());
	sort(right.begin(), right.end());

    int sum = 0;

    for (int i = 0; i < left.size(); i++) {
        // answer for part 1
		// sum += abs(left[i] - right[i]);

        // answer for part 2
		sum += left[i] * count(right.begin(), right.end(), left[i]);
	}

    return sum;
}

int main() {
	cout << problem1() << "\n";
	return 0;
}
