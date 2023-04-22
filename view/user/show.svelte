<script>
  export let user = {}
</script>

<h1>Show User</h1>

<table border="1" cellpadding="10">
  <thead>
    {#each Object.keys(user) as key}
      <th>{key}</th>
    {/each}
  </thead>
  <tr>
    {#each Object.keys(user) as key}
      <td>{user[key]}</td>
    {/each}
  </tr>
</table>

<br />

<a href={`/user`}>Back</a>
<span> | </span>
<a href={`/user/${user.id || 0}/edit`}>Edit</a>

<style>
  table {
    border-collapse: collapse;
  }
</style>
