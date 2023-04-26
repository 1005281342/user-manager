<script>
  export let user;
</script>

<style>
  /* 表单样式 */
  form {
    display: flex;
    flex-direction: column;
    align-items: center;
    margin-top: 50px;
  }

  .form-group {
    margin-bottom: 20px;
    width: 100%;
    max-width: 500px;
  }

  label {
    display: block;
    font-weight: bold;
    margin-bottom: 5px;
  }

  .form-control {
    display: block;
    width: 100%;
    padding: 10px;
    font-size: 16px;
    border: 1px solid #ccc;
    border-radius: 5px;
    box-sizing: border-box;
  }

  /* 按钮样式 */
  .btn {
    display: inline-block;
    padding: 10px 20px;
    font-size: 16px;
    font-weight: bold;
    text-align: center;
    text-decoration: none;
    border-radius: 5px;
    cursor: pointer;
    margin-right: 10px;
  }

  .btn-primary {
    color: #fff;
    background-color: #4caf50;
    border-color: #4caf50;
  }

  .btn-secondary {
    color: #fff;
    background-color: #6c757d;
    border-color: #6c757d;
  }

  .btn-info {
    color: #fff;
    background-color: #17a2b8;
    border-color: #17a2b8;
  }

  .btn:hover,
  .btn:focus {
    text-decoration: none;
  }

  .btn-primary:hover,
  .btn-primary:focus {
    background-color: #388e3c;
    border-color: #388e3c;
  }

  .btn-secondary:hover,
  .btn-secondary:focus {
    background-color: #5a6268;
    border-color: #545b62;
  }

  .btn-info:hover,
  .btn-info:focus {
    background-color: #117a8b;
    border-color: #10707f;
  }

  /* 返回链接样式 */
  .back-link {
    margin-top: 20px;
    text-align: center;
  }

  .back-link a {
    color: #007bff;
    text-decoration: none;
  }

  .back-link a:hover {
    text-decoration: underline;
  }
</style>

<div style="display:flex;justify-content:center;">
  <div class="container">
    <h1 style="text-align:center;">Edit User {user.ID}</h1>

    <form method="post" action={`/user/${user.ID || 0}`}>
      <input type="hidden" name="_method" value="patch" />

      <div class="form-group">
        <label for="email">Email:</label>
        <input type="email" name="email" id="email" value="{user.email}" required class="form-control" />
      </div>

      <div class="form-group">
        <label for="password">Password:</label>
        <input type="password" name="password" id="password" value="" required class="form-control" />
      </div>

      <div style="display:flex;justify-content:center;">
        <button type="submit" class="btn btn-primary" style="background-color:#4caf50;">Update User</button>
        <a href={`/user/${user.ID || 0}`} class="btn btn-info">Show User</a>
        <a href="/user" class="btn btn-secondary">Back</a>
      </div>
    </form>

    <div class="back-link">
      <a href={`/user/${user.ID || 0}`}>Back to User</a>
    </div>
  </div>
</div>
