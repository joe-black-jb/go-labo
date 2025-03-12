<h1>ログイン</h1>
<form method="POST">
  <div>
    <label>
      Email
      <input name="email" type="email" class="input-box">
    </label>
  </div>
  <div>
    <label>
      Password
      <input name="password" type="password" class="input-box">
    </label>
  </div>
  <div>
    <button class="login-button">Log in</button>
  </div>
</form>

<style>
  .input-box {
    border: 2px solid;
  }
  .login-button {
    background-color: #4CAF50;
    color: white;
    padding: 14px 20px;
    margin: 8px 0;
    border: none;
    cursor: pointer;
    /* width: 100%; */
    border-radius: 5%;
  }
  .login-button:hover {
    background-color: #45a049; /* ホバー時の背景色 */
  }
</style>