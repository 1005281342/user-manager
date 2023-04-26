<script>
  export let user = {}
</script>

<style>
  /* 用户详情样式 */
  .user-details {
    display: flex;
    flex-direction: column;
    align-items: center;
    margin-top: 50px;
  }

  .user-info {
    display: flex;
    flex-direction: column;
    align-items: center;
    margin-top: 20px;
    width: 100%;
    max-width: 500px;
  }

  .user-email {
    font-size: 1.2rem;
    margin-bottom: 0.5rem;
    display: flex;
    align-items: center;
  }

  .user-email span {
    font-weight: bold;
    margin-right: 0.5rem;
  }

  /* 按钮样式 */
  .actions {
    display: flex;
    align-items: center;
    margin-top: 20px;
  }

  .button {
    padding: 10px 20px;
    background-color: #4caf50;
    color: white;
    border: none;
    border-radius: 5px;
    font-size: 1rem;
    margin-right: 10px;
    cursor: pointer;
  }

  .button:hover,
  .button:active {
    background-color: #3e8e41;
  }

  .back-button,
  .button {
    display: flex;
    align-items: center;
    text-decoration: none;
  }

  .back-button {
    padding: 10px 20px;
    border-radius: 5px;
    color: #fff;
    background-color: #ccc;
    margin-right: 10px;
  }

  .back-button:hover {
    background-color: #4caf50;
  }
</style>

<div style="display:flex;justify-content:center;">
  <div class="user-details">
    <h1 style="text-align:center;">User Details</h1>

    <div class="user-info">
      <div class="user-email">
        <span>Email: </span>{user.email}
      </div>
    </div>

    <div class="actions">
      <a href="/user" class="back-button">Back</a>
      <a href={`/user/${user.ID}/edit`} class="button">Edit User</a>
    </div>
  </div>
</div>
