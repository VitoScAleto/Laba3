#include <iostream>
#include <vector>

using namespace std;

template <typename T>
ostream& operator<<(ostream& os, const vector<T>& nums) // возвращаем ссылку на поток что бы cout мог ухватиться и вывести что-то следом
{

    for (auto i : nums)
    {
        os << i << " "; // os так как необязательно поток вывода консольный
    }

    return os;
}

template <typename T>
vector<T> operator+(const vector<T>& nums1, const vector<T>& nums2)
{

    vector<T> result;

    for (size_t i = 0; i < nums1.size(); i++)
    {
        result.push_back(nums1[i] + nums2[i]);
    }

    return result;
}

int main()
{

    vector<vector<int>> numbers1 = { {-10, 2, -6, 20}, {2, 0} };
    vector<vector<int>> numbers2 = { {-10, 5, -3, 12}, {2, 0} };
    cout << numbers1 + numbers2 << " ";

    return 0;
}
