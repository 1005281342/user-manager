<script>
  export let user = {}
</script>

<h1>Edit User</h1>

<form method="post" action={`/user/${user.id || 0}`}>
  <input type="hidden" name="_method" value="patch" />
  <!-- Add input fields here -->
  <input type="submit" value="Update User" />
</form>

<br />

<a href={`/user`}>Back</a>
<span> | </span>
<a href={`/user/${user.id || 0}`}>Show User</a>
