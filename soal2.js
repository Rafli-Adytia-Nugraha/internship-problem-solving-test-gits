function denseRanking(scores, gits) {
    const unique = [];
    for (let i = 0; i < scores.length; i++) {
      if (i === 0 || scores[i] !== scores[i - 1]) {
        unique.push(scores[i]);
      }
    }

    const n = unique.length;
    const result = [];
    let i = n - 1;

    for (const g of gits) {
      while (i >= 0 && g >= unique[i]) {
        i--;
      }
      result.push(i + 2);
    }

    return result;
}

function printFormatted(scores, gits) {
  console.log("Sampel Input:");
  console.log(scores.length);
  console.log(scores.join(" "));
  console.log(gits.length);
  console.log(gits.join(" "));
  console.log("\nSampel Output:");
  console.log(denseRanking(scores, gits).join(" "));
  console.log("\n");
}

const scores1 = [100, 100, 50, 40, 40, 20, 10];
const gits1 = [5, 25, 50, 120];
printFormatted(scores1, gits1);

const scores2 = [90, 80, 80, 75, 60];
const gits2 = [50, 65, 77, 90, 102];
printFormatted(scores2, gits2);

const scores3 = [200, 150, 150, 100, 90];
const gits3 = [70, 120, 150, 210];
printFormatted(scores3, gits3);