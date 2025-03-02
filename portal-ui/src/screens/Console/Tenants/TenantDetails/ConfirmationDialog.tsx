import React, { useState } from "react";
import { Theme } from "@mui/material/styles";
import { Button } from "mds";
import createStyles from "@mui/styles/createStyles";
import withStyles from "@mui/styles/withStyles";
import {
  containerForHeader,
  tenantDetailsStyles,
} from "../../Common/FormComponents/common/styleLibrary";
import {
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle,
  LinearProgress,
} from "@mui/material";

interface IConfirmationDialog {
  classes: any;
  open: boolean;
  cancelLabel: string;
  okLabel: string;
  onClose: any;
  cancelOnClick: any;
  okOnClick: any;
  title: string;
  description: string;
}

const styles = (theme: Theme) =>
  createStyles({
    ...tenantDetailsStyles,
    ...containerForHeader(theme.spacing(4)),
  });

const ConfirmationDialog = ({
  classes,
  open,
  cancelLabel,
  okLabel,
  onClose,
  cancelOnClick,
  okOnClick,
  title,
  description,
}: IConfirmationDialog) => {
  const [isSending, setIsSending] = useState<boolean>(false);
  const onClick = () => {
    setIsSending(true);
    if (okOnClick !== null) {
      okOnClick();
    }
    setIsSending(false);
  };
  if (!open) return null;
  return (
    <Dialog
      open={open}
      onClose={onClose}
      aria-labelledby="alert-dialog-title"
      aria-describedby="alert-dialog-description"
    >
      <DialogTitle id="alert-dialog-title">{title}</DialogTitle>
      <DialogContent>
        {isSending && <LinearProgress />}
        <DialogContentText id="alert-dialog-description">
          {description}
        </DialogContentText>
      </DialogContent>
      <DialogActions>
        <Button
          id="cancel-dialog"
          variant={"callAction"}
          onClick={cancelOnClick}
          disabled={isSending}
          label={cancelLabel || "Cancel"}
        />
        <Button
          id="accept-dialog"
          variant={"secondary"}
          onClick={onClick}
          label={okLabel || "Ok"}
        />
      </DialogActions>
    </Dialog>
  );
};

export default withStyles(styles)(ConfirmationDialog);
