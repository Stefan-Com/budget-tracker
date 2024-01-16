export async function register(info: User) {
  return new Promise(async (resolve, reject) => {
    const body = JSON.stringify({
      ...info,
      balance: parseFloat(String(info.balance)),
    });
    const response = await fetch("http://localhost:8000/register", {
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