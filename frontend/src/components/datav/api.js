function callLog (params) {
  return fetch('http://127.0.0.1:8080/api/calllog').then((res) => {
    return res.json().then((data) => {
      return data
    })
  })
}

function systeminfo (params) {
  return fetch('http://127.0.0.1:8080/api/systeminfo').then((res) => {
    return res.json().then((data) => {
      return data
    })
  })
}

function queueStatus (params) {
  return fetch('http://127.0.0.1:8080/api/queuestatus').then((res) => {
    return res.json().then((data) => {
      return data
    })
  })
}

function extensionStatus (params) {
  return fetch('http://127.0.0.1:8080/api/extensionstatus').then((res) => {
    return res.json().then((data) => {
      return data
    })
  })
}

function trunkStatus (params) {
  return fetch('http://127.0.0.1:8080/api/trunkstatus').then((res) => {
    return res.json().then((data) => {
      return data
    })
  })
}

export { callLog, systeminfo, queueStatus, extensionStatus, trunkStatus }
