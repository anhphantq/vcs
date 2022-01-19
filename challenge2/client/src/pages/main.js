import AccountInfo from "../component/AccountInfo";
import AdminPanel from "../component/AdminPanel";
import { Grid, Button } from "@mui/material";
import { useSelector } from "react-redux";
import SignIn from "./signin";
import { useDispatch } from "react-redux";
import { logout } from "../redux/account-slice";
import { deleteList } from "../redux/admin-slice";

function Main() {
  const dp = useDispatch();
  const isLoggedIn = useSelector((state) => state.accountSlice.isLoggedIn);
  if (!isLoggedIn) return <SignIn></SignIn>;
  const hdLogout = () => {
    dp(deleteList());
    dp(logout());
  };

  return (
    <Grid container>
      <Grid item xs={12}>
        <AccountInfo></AccountInfo>
        <Button onClick={hdLogout}>LOGOUT</Button>
      </Grid>
      <Grid item xs={12}>
        <AdminPanel></AdminPanel>
      </Grid>
    </Grid>
  );
}

export default Main;
