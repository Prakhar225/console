// This file is part of MinIO Console Server
// Copyright (c) 2022 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

import React, { Fragment, useEffect, useState } from "react";
import { Theme } from "@mui/material/styles";
import { useSelector } from "react-redux";
import Grid from "@mui/material/Grid";
import createStyles from "@mui/styles/createStyles";
import withStyles from "@mui/styles/withStyles";
import { AppState, useAppDispatch } from "../../../../store";

import { CircleIcon, ObjectManagerIcon } from "../../../../icons";
import { Box } from "@mui/material";
import { toggleList } from "../../ObjectBrowser/objectBrowserSlice";
import { selFeatures } from "../../consoleSlice";
import { selDirectPVMode, selOpMode } from "../../../../systemSlice";
import { ApplicationLogo, Button } from "mds";

const styles = (theme: Theme) =>
  createStyles({
    headerContainer: {
      width: "100%",
      minHeight: 83,
      display: "flex",
      backgroundColor: "#fff",
      left: 0,
      borderBottom: "1px solid #E5E5E5",
    },
    label: {
      display: "flex",
      justifyContent: "flex-start",
      alignItems: "center",
    },
    rightMenu: {
      display: "flex",
      justifyContent: "flex-end",
      paddingRight: 20,
      "& button": {
        marginLeft: 8,
      },
    },
    logo: {
      marginLeft: 34,
      "& svg": {
        width: 150,
      },
    },
    middleComponent: {
      display: "flex",
      justifyContent: "center",
      alignItems: "center",
    },
    indicator: {
      position: "absolute",
      display: "block",
      width: 15,
      height: 15,
      top: 0,
      right: 4,
      marginTop: 4,
      transitionDuration: "0.2s",
      color: "#32C787",
      "& svg": {
        width: 10,
        height: 10,
        top: "50%",
        left: "50%",
        transitionDuration: "0.2s",
      },
      "&.newItem": {
        color: "#2781B0",
        "& svg": {
          width: 15,
          height: 15,
        },
      },
    },
  });

interface IPageHeader {
  classes: any;
  label: any;
  actions?: any;
  middleComponent?: React.ReactNode;
}

const PageHeader = ({
  classes,
  label,
  actions,
  middleComponent,
}: IPageHeader) => {
  const dispatch = useAppDispatch();

  const sidebarOpen = useSelector(
    (state: AppState) => state.system.sidebarOpen
  );
  const operatorMode = useSelector(selOpMode);
  const directPVMode = useSelector(selDirectPVMode);
  const managerObjects = useSelector(
    (state: AppState) => state.objectBrowser.objectManager.objectsToManage
  );
  const features = useSelector(selFeatures);
  const managerOpen = useSelector(
    (state: AppState) => state.objectBrowser.objectManager.managerOpen
  );
  const newItems = useSelector(
    (state: AppState) => state.objectBrowser.objectManager.newItems
  );
  const licenseInfo = useSelector(
    (state: AppState) => state?.system?.licenseInfo
  );

  const [newObject, setNewObject] = useState<boolean>(false);

  const { plan = "" } = licenseInfo || {};

  let logoPlan = "AGPL";

  if (plan === "STANDARD" || plan === "ENTERPRISE") {
    logoPlan = plan.toLowerCase();
  }

  useEffect(() => {
    if (managerObjects.length > 0 && !managerOpen) {
      setNewObject(true);
      setTimeout(() => {
        setNewObject(false);
      }, 300);
    }
  }, [managerObjects.length, managerOpen]);

  if (features.includes("hide-menu")) {
    return <Fragment />;
  }
  return (
    <Grid
      container
      className={`${classes.headerContainer} page-header`}
      direction="row"
      alignItems="center"
    >
      <Grid
        item
        xs={12}
        sm={12}
        md={middleComponent ? 4 : 6}
        className={classes.label}
        sx={{
          paddingTop: ["15px", "15px", "0", "0"],
        }}
      >
        {!sidebarOpen && (
          <div className={classes.logo}>
            {!operatorMode && !directPVMode ? (
              <ApplicationLogo
                applicationName={"console"}
                subVariant={
                  logoPlan as
                    | "AGPL"
                    | "simple"
                    | "standard"
                    | "enterprise"
                    | undefined
                }
              />
            ) : (
              <Fragment>
                {directPVMode ? (
                  <ApplicationLogo applicationName={"directpv"} />
                ) : (
                  <ApplicationLogo applicationName={"operator"} />
                )}
              </Fragment>
            )}
          </div>
        )}
        <Box
          sx={{
            color: "#000",
            fontSize: 18,
            fontWeight: 700,
            marginLeft: "21px",
            display: "flex",
          }}
        >
          {label}
        </Box>
      </Grid>
      {middleComponent && (
        <Grid
          item
          xs={12}
          sm={12}
          md={4}
          className={classes.middleComponent}
          sx={{ marginTop: ["10px", "10px", "0", "0"] }}
        >
          {middleComponent}
        </Grid>
      )}
      <Grid
        item
        xs={12}
        sm={12}
        md={middleComponent ? 4 : 6}
        className={classes.rightMenu}
      >
        {actions && actions}
        {managerObjects && managerObjects.length > 0 && (
          <Button
            aria-label="Refresh List"
            onClick={() => {
              dispatch(toggleList());
            }}
            icon={
              <Fragment>
                <div
                  className={`${classes.indicator} ${
                    newObject ? "newItem" : ""
                  }`}
                  style={{
                    opacity: managerObjects.length > 0 && newItems ? 1 : 0,
                  }}
                >
                  <CircleIcon />
                </div>
                <ObjectManagerIcon
                  style={{ width: 20, height: 20, marginTop: -2 }}
                />
              </Fragment>
            }
            id="object-manager-toggle"
            style={{ position: "relative", padding: "0 10px" }}
          />
        )}
      </Grid>
    </Grid>
  );
};

export default withStyles(styles)(PageHeader);
