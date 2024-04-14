#include <iostream>
#include <cmath>
#include <vector>
#include <locale.h>
#include <iomanip>

using namespace std;
void Zadanie1_21(void);
int cofe(double Tsr, double Tc, double r, int t, vector<double>& Cofe);
double aproxA(vector<double> Cofe);
double aproxB(vector<double> Cofe, double a);
double Correl(vector<double> Cofe);
void OutPutZadanie3(void);
int main()
{
	setlocale(LC_ALL, "Rus");
	Zadanie1_21();
	OutPutZadanie3();













}

void Zadanie1_21(void)
{
	float x = 0.0;
	float y = sqrt(-pow(x+2,2)+pow(2,2));// функция окружности
	float yLog = log(x) / x;
	cout << "Полуокружность:\n";
	for(;x>-4.0 ; x -= 0.1)
	{
		y = sqrt(-pow(x + 2, 2) + pow(2, 2));
		cout << " y = " <<setprecision(4) <<round(y*1000)/1000 << "\t x = " <<x<< "\t dx = "<< 0.1 <<endl;
	}
	cout << "Натуральный логарифм на х:\n";
	x = 0.5;
	for (; x < 2; x += 0.1)
	{
		yLog = log(x) / x;
		cout << "y = " << round(yLog * 1000) / 1000 << "\tx =" << x << "\tdx =" << 0.1 << endl;
	}



}
int cofe(double Tsr, double Tc, double r, int t, vector<double>& Cofe) {
	for (int i = 0; i <= t; i++) {
		Cofe.push_back(Tc);        //добавляем изначальное значение
		Tc = Tc - r * (Tc - Tsr);  //высчитываем новое по формуле
	}
	return 0;
}

double aproxA(vector<double> Cofe) {

	double ET = 0, Et = 0, ETt = 0, Et2 = 0;
	int len = Cofe.size();  //количество измерений

	for (int i = 0; i < len; i++) {
		ET += Cofe[i];                //сумма по температуре
		Et += i;                      //сумма по времени

		ETt += Cofe[i] * i;           //сумма для произведения по оси температуры и времени
		Et2 += i * i;                   //сумма для квадрата температуры
	}
	return (len * ETt - (Et * ET)) / (len * Et2 - Et * Et);
}

double aproxB(vector<double> Cofe, double a) {

	double ET = 0, Et = 0;
	int len = Cofe.size();  //количество измерений

	for (int i = 0; i < len; i++) {
		ET += Cofe[i];              //сумма по температуре
		Et += i;                    //сумма по времени
	}
	return (ET - a * Et) / len;
}

double Correl(vector<double> Cofe) {

	double sumTemper = 0;   //сумма температур
	for (double T : Cofe) {
		sumTemper += T;
	}

	int len = Cofe.size();  //количество измерений
	double TMedium = sumTemper / len;   //среднее значение всех измеренных значений
	double tMedium = (len - 1) * len / 2;
	double sumNumbers = 0;   //арифметическая сумма значений произведения температуры и времени
	double tSumSquare = 0;   //арифметическая сумма квадрата времени
	double TSumSquare = 0;   //арифметическая сумма квадрата температуры

	for (int i = 0; i < len; i++) {
		sumNumbers += ((i - tMedium) * (Cofe[i] - TMedium));
		tSumSquare += ((i - tMedium) * (i - tMedium));
		TSumSquare += ((Cofe[i] - TMedium) * (Cofe[i] - TMedium));
	}
	return sumNumbers / sqrt(TSumSquare * tSumSquare);
}

void OutPutZadanie3(void)
{
	int t;            //время
	double Tsr, Tc, r;   //температура среды, температура кофе, коэф.остывания
	cout << "Введите температуру среды:" << endl;
	cin >> Tsr;
	cout << "Введите температуру кофе:" << endl;
	cin >> Tc;
	cout << "Введите коэффициент остывания:" << endl;
	cin >> r;
	cout << "Введите время остывания:" << endl;
	cin >> t;

	vector<double> Cofe;
	cofe(Tsr, Tc, r, t, Cofe);  //расчитываем температуру по времени

	cout << " time  | coffee temperature" << endl;
	int time = 0;
	for (auto temperature : Cofe) {
		cout << "-----------------------" << endl;
		cout << setw(4) << time << "   |" << setw(17) << temperature << endl;  //вывели все значение от 0 до времени
		time++;
	}

	double a = aproxA(Cofe);  // в a помещаем отклонение по оси времени
	double b = aproxB(Cofe, a);   //в b помещаем отклонение по оси температуры
	cout << endl << "Линия апроксимации: " << "T = " << a << " * t + " << b << endl; //линия апроксимации

	double correl = Correl(Cofe); //высчитываем коэф.корреляции
	cout << endl << "Погрешность измерений " << correl << endl;  //погрешность измерений, нужно для вычисления погрешности


}