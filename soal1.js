function A000124Formatted(n) {
    if (n <= 0) return "";

    const parts = [];
    for (let k = 0; k < n; k++) {
        const val = (k * (k + 1)) / 2 + 1;
        parts.push(val.toString());
    }

    return parts.join("-");
}

const inputs = [7, 1, 12];

for (const n of inputs) {
    console.log(`Input: ${n}`);
    console.log(`Output: ${A000124Formatted(n)}\n`);
}