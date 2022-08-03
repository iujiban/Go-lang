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
  const messageBuffer = new Buffer.from(message)
  const secrets = [...messageBuffer]
  const polynomial = new Array(K).fill(0)


  polynomial[0] = secrets[0]

  for(let i = 1; i< K ; i++) {
    polynomial[i] = getRandomInt()
  }

  const shares = new Array(N)

  for(let i = 0; i< N; i++) {
    shares[i] = new Array(2)
    do {
      x = getRandomInt();


    } while (IsExistX(shares, i, x));

    shares[i][0] = x
    shares[i][1] =  evaludatePolynomial(polynomial, x)

    return shares

  }
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
  let secret = 0
  
  for (let i = 0 ; i< shares.length; i++) {
    const share = shares[i]
    const x = shares[0]
    const y = shares[1]

    let numerator = 1
    let denominator = 1
    
    for (let j = 0; j < shares.length; j++) {
      if(i != j) {
        numerator = numerator * (-shares[j][0])
        numerator = modulo(numerator, primeNumber)

        denominator = denominator * (x - shares[j][0])
        denominator = modulo(denominator, primeNumber)



      }
    }
    inversed  = modInverse(denominator, primeNumber)

    secret = secret + y * (numerator * inversed)
    secret = modulo(secret, primeNumber)


  }

  var buffer = new Buffer.alloc(1);
  buffer[0] = secret
  return buffer

}

const shares = create('1', 5, 10)
console.log(combine(shares.slice(1,6)).toString())