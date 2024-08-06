// let apiUrl = 'http://127.0.0.1:8080/'
let apiUrl = '/'

function callLog (params) {
  return fetch(`${apiUrl}/api/calllog`).then((res) => {
    return res.json().then((data) => {
      return data
    })
  })
}

function systeminfo (params) {
  return fetch(`${apiUrl}/api/systeminfo`).then((res) => {
    return res.json().then((data) => {
      return data
    })
  })
}

function queueStatus (params) {
  return fetch(`${apiUrl}/api/queuestatus`).then((res) => {
    return res.json().then((data) => {
      return data
    })
  })
}

function extensionStatus (params) {
  return fetch(`${apiUrl}/api/extensionstatus`).then((res) => {
    return res.json().then((data) => {
      return data
    })
  })
}

function trunkStatus (params) {
  return fetch(`${apiUrl}/api/trunkstatus`).then((res) => {
    return res.json().then((data) => {
      return data
    })
  })
}

function networkInfo (params) {
  return fetch(`${apiUrl}/api/networkinfo`).then((res) => {
    return res.json().then((data) => {
      return data
    })
  })
}

export { callLog, systeminfo, queueStatus, extensionStatus, trunkStatus, networkInfo }
