function flipStr(string = 'ABCD123') {
    const strArr = string.split('');
    let flipArr = []
    strArr.map((str) => {
      let n =  str.charCodeAt()
      let z = 'z'.charCodeAt()
        n = n+1
      if(n > z) {
        n = 'a'.charCodeAt();
        flipArr.push(String.fromCharCode(n))
      } 
      else flipArr.push(String.fromCharCode(n))
    })
    return flipArr.join('');
}

const result = flipStr();
console.log(result) // bcdza234