<script>
  let email = '';
  let password = '';

  const handleSubmit = async (event) => {
    event.preventDefault();

    const response = await fetch('/user', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ email, password })
    });

    if (response.ok) {
      window.location.href = '/user';
    } else {
      alert('Failed to create user');
    }
  };
</script>

<h1 style="text-align:center;">New User</h1>

<div style="display:flex;justify-content:center;">
  <form on:submit={handleSubmit}>
    <div style="display:flex;flex-direction:column;align-items:center;width:25rem;margin:auto;">
      <label for="email">Email:</label>
      <input type="email" id="email" bind:value={email} required />
      <br />
      <label for="password">Password:</label>
      <input type="password" id="password" bind:value={password} required />
      <br />
      <div style="display:flex;justify-content:space-between;align-items:center;margin-top:1rem;width:20rem;">
        <a href="/user" class="back-button" style="background-color:#ccc;text-decoration:none;padding:0.5rem 1rem;border-radius:5px;color:#fff;height:2.5rem;display:flex;align-items:center;">Back</a>
        <button type="submit" class="submit-button" style="background-color:#4caf50;border:none;justify-content:center;width:10rem;margin-left:auto;margin-right:auto;padding:0.5rem 1rem;border-radius:5px;color:#fff;height:2.5rem;display:flex;align-items:center;">Create User</button>
      </div>
    </div>
  </form>
</div>

<style>
  input {
    margin: 0.5rem;
    padding: 0.5rem;
    border: 1px solid #ccc;
    border-radius: 0.25rem;
    width: 20rem;
  }

  h1 {
    text-align: center;
  }
</style>
