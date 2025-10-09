function replaceAt(s, i, r) {
    return s.slice(0, i) + r + s.slice(i + 1);
}

function makePalindrome(s, left, right, k) {
    if (left >= right) return [s, k, true];
    if (s[left] === s[right]) return makePalindrome(s, left + 1, right - 1, k);
    if (k <= 0) return [s, k, false];

    const bigger = s[left] > s[right] ? s[left] : s[right];
    s = replaceAt(s, left, bigger);
    s = replaceAt(s, right, bigger);
    return makePalindrome(s, left + 1, right - 1, k - 1);
}

function maximizePalindrome(s, left, right, k) {
    if (left > right || k <= 0) return s;
    if (s[left] === '9' && s[right] === '9')
    return maximizePalindrome(s, left + 1, right - 1, k);

    if (left === right && k > 0)
    return replaceAt(s, left, '9');

    if (s[left] !== '9' || s[right] !== '9') {
        let need = 0;
        if (s[left] !== '9') need++;
        if (s[right] !== '9') need++;
        if (k >= need) {
            s = replaceAt(s, left, '9');
            s = replaceAt(s, right, '9');
            return maximizePalindrome(s, left + 1, right - 1, k - need);
        }
    }

    return maximizePalindrome(s, left + 1, right - 1, k);
}

function highestPalindrome(s, k) {
    const [pal, rem, ok] = makePalindrome(s, 0, s.length - 1, k);
    if (!ok) return "-1";
    return maximizePalindrome(pal, 0, pal.length - 1, rem);
}

const input1 = "3943", k1 = 1;
console.log(`Input: ${input1}\nk: ${k1}\nOutput: ${highestPalindrome(input1, k1)}\n`);

const input2 = "932239", k2 = 2;
console.log(`Input: ${input2}\nk: ${k2}\nOutput: ${highestPalindrome(input2, k2)}\n`);

const input3 = "12321", k3 = 1;
console.log(`Input: ${input3}\nk: ${k3}\nOutput: ${highestPalindrome(input3, k3)}\n`);