const primeNumber = 95971

function modulo(a, m) {
  result = a % m
  return result < 0 ? result + m : result
}

function modInverse(a, m) {
  [a, m] = [Number(a), Number(m)]
  
  if (Number.isNaN(a) || Number.isNaN(m)) {
    return NaN
  }

  a = (a % m + m) % m
  if (!a || m < 2) {
    return NaN
  }

  const s = []
  let b = m
  while(b) {
    [a, b] = [b, a % b]
    s.push({a, b})
  }
  if (a !== 1) {
    return NaN
  }

  let x = 1, y = 0
  for(let i = s.length - 2; i >= 0; --i) {
    [x, y] = [y,  x - y * Math.floor(s[i].a / s[i].b)]
  }
  return (y % m + m) % m
}

function getRandomInt() {
  return Math.floor((Math.random() * primeNumber) + 1) % primeNumber
}

function IsExistX(shares, N, x) {
  for(let i = 0 ; i < N ; i++) {
    if(shares[i][0] == x) {
      return true
    }
  }

  return false
}

function create(message, K, N) {
}

function evaludatePolynomial(polynomial, x) {  
  const last = polynomial.length - 1
  let result = polynomial[last]

  for (let i = last - 1; i >= 0; i--) {
      result = result * x
      result = result + polynomial[i]
      result = modulo(result, primeNumber)
  }

  return result
}

function combine(shares) {
}

const shares = create('1', 5, 10)
console.log(combine(shares.slice(1,6)).toString())