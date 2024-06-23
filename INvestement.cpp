#include<bits/stdc++.h>
using namespace std;
int main(){
    int paycheck;
    cout<<"enter ur salary";
    cin>>paycheck;
    int responsibilty;
    cout<<"amount of moneys goes to ur parent or in donation";
    cin>>responsibilty;
    int leftover=paycheck-responsibilty;
    int expense;
    cout<<"total of expense u made monthly (loan,rent,bills)";
    cin>>expense;
    if(leftover/1.8>=expense)
    cout<<"ur fixed expense are on control ";
    else
    cout<<"u should limit expense";
    
}