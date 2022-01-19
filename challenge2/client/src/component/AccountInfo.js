import { Grid } from "@mui/material";
import { getInfo } from "../redux/account-slice";
import { useDispatch, useSelector } from "react-redux";

function AccountInfo() {
  const dp = useDispatch();

  dp(getInfo());

  const info = useSelector((state) => state.accountSlice.info);

  return (
    info && (
      <Grid container>
        <Grid xs={12}>YOUR ACCOUNT</Grid>
        <Grid xs={12}>Your ID: {`${info.user_id}`} </Grid>
        <Grid xs={12}>Your Name: {`${info.username}`} </Grid>
        <Grid xs={12}>Your Email: {`${info.email}`} </Grid>
        <Grid xs={12}>Your Role ID: {`${info.role_id}`} </Grid>
      </Grid>
    )
  );
}

export default AccountInfo;
