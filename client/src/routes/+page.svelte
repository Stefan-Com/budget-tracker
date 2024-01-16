<script>
  async function logout() {
    return new Promise(async (resolve, reject) => {
      const response = await fetch("http://localhost:8000/logout", {
        method: "DELETE",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
      }).then((res) => res.json());
      console.log(response.response);
      resolve(null);
    });
  }
</script>

<div>
  <h1>This is the homepage</h1>
  <button on:click={logout}>Logout</button><br />
</div>
