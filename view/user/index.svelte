<script>
  export let users = []
</script>

<h1>User Index</h1>

<table border="1" cellpadding="10">
  {#if users.length > 0}
    <thead>
      {#each Object.keys(users[0]) as key}
        <th>{key}</th>
      {/each}
    </thead>
  {/if}
  {#each users as user}
    <tr>
      {#each Object.keys(user) as key}
        {#if key.toLowerCase() === "id"}
          <td><a href={`/user/${user.id || 0}`}>{user[key]}</a></td>
        {:else}
          <td>{user[key]}</td>
        {/if}
      {/each}
    </tr>
  {/each}
</table>

<br />

<a href={`/user/new`}>New User</a>

<style>
  table {
    border-collapse: collapse;
  }
</style>
