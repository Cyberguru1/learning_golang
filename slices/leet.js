#!/usr/bin/node



var digits = [9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9]

len = digits.length
check = digits[len -1]

var i = 0
if (check == 9) {
    while (digits[len - i - 1] == 9) {
        i++;
    }
}

if (i == 0) {
    last_num = digits[len-1]
    last_num += 1
    digits.pop()
    digits.push(last_num)
}
else if (i == len){
    last_nums = digits.slice(len-i, len)
    last_num = last_nums.reduce((t, x) => {return t + x.toString()},)
    console.log(last_num)
    last_num =  (parseInt(last_num).toString() == '1e+99') ? BigInt(last_num, 10) + 1n: parseInt(last_num) + 1;
    console.log(last_num)
    digits = last_num.toString().split("").map(x => parseInt(x))

}
else {
    last_nums = digits.slice(len-i, len)
    last_num = last_nums.reduce((t, x) => {return t + x.toString()},)
    last_num = parseInt(last_num)
    last_num += 1
    res = last_num.toString().split("").map(x => parseInt(x))
    digits[len - res.length] += 1
    digits = digits.slice(0, len - res.length+1).concat(res.slice(1, res.length))
}
console.log(digits)