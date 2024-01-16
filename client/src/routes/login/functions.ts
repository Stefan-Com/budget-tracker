export async function login(info: any) {
  return new Promise(async (resolve, reject) => {
    const body = JSON.stringify(info);
    const response = await fetch("http://localhost:8000/login", {
      method: "POST",
      body: body,
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    }).then((res) => res.json());
    console.log(response.response);
    resolve(null);
  });
}