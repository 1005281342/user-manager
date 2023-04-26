<script>
  export let users = [];
</script>

<style>
  /* 表格样式 */
  table {
    width: 100%;
    border-collapse: collapse;
    margin-bottom: 1rem;
    background-color: #f8f8f8;
    box-shadow: 0 0 20px rgba(0, 0, 0, 0.1);
    font-size: 0.9rem;
    font-family: sans-serif;
  }

  th,
  td {
    padding: 0.5rem;
    text-align: left;
    vertical-align: middle;
  }

  th {
    background-color: #e0e0e0;
    font-weight: bold;
    text-transform: uppercase;
    letter-spacing: 0.1em;
    border-bottom: 2px solid #bdbdbd;
  }

  td {
    border-bottom: 1px solid #bdbdbd;
  }

  /* 新用户链接样式 */
  a.new-user {
    display: inline-block;
    background-color: #4caf50;
    color: #fff;
    padding: 0.5rem 1rem;
    border-radius: 4px;
    text-decoration: none;
    transition: background-color 0.2s;
  }

  a.new-user:hover {
    background-color: #388e3c;
  }

  /* 隐藏密码列 */
  td.password,
  th.password {
    display: none;
  }

  /* 用户ID链接样式 */
  td.id a,
  th.id {
    color: #2196f3;
    text-decoration: none;
  }

  td.id a:hover {
    text-decoration: underline;
  }
</style>

<h1 style="text-align:center;">User Index</h1>

<div style="display:flex;justify-content:center;">
  {#if users.length > 0}
    <table>
      <thead>
      <tr>
        {#each Object.keys(users[0]) as key}
          {#if key.toLowerCase() === "password"}
            <th class="password">{key}</th>
          {:else if key.toLowerCase() === "id"}
            <th class="id">{key}</th>
          {:else}
            <th>{key}</th>
          {/if}
        {/each}
      </tr>
      </thead>
      <tbody>
      {#each users as user}
        <tr>
          {#each Object.keys(user) as key}
            {#if key.toLowerCase() === "password"}
              <td class="password">{user[key]}</td>
            {:else if key.toLowerCase() === "id"}
              <td class="id"><a href={`/user/${user[key] || 0}`}>{user[key]}</a></td>
            {:else}
              <td>{user[key]}</td>
            {/if}
          {/each}
        </tr>
      {/each}
      </tbody>
    </table>
  {/if}
</div>

<div style="display:flex;justify-content:center;">
  <a href="/user/new" class="new-user" style="background-color:#4caf50;">New User</a>
</div>