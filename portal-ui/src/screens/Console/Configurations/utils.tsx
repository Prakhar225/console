// This file is part of MinIO Console Server
// Copyright (c) 2021 MinIO, Inc.
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
import React from "react";
import PublicIcon from "@mui/icons-material/Public";
import CompressIcon from "@mui/icons-material/Compress";
import CodeIcon from "@mui/icons-material/Code";
import LocalHospitalIcon from "@mui/icons-material/LocalHospital";
import FindReplaceIcon from "@mui/icons-material/FindReplace";
import VpnKeyIcon from "@mui/icons-material/VpnKey";
import PendingActionsIcon from "@mui/icons-material/PendingActions";
import CallToActionIcon from "@mui/icons-material/CallToAction";
import { IElement, IElementValue } from "./types";

export const configurationElements: IElement[] = [
  {
    icon: <PublicIcon />,
    configuration_id: "region",
    configuration_label: "Region",
  },
  {
    icon: <CompressIcon />,
    configuration_id: "compression",
    configuration_label: "Compression",
  },
  {
    icon: <CodeIcon />,
    configuration_id: "api",
    configuration_label: "API",
  },
  {
    icon: <LocalHospitalIcon />,
    configuration_id: "heal",
    configuration_label: "Heal",
  },
  {
    icon: <FindReplaceIcon />,
    configuration_id: "scanner",
    configuration_label: "Scanner",
  },
  {
    icon: <VpnKeyIcon />,
    configuration_id: "etcd",
    configuration_label: "Etcd",
  },
  {
    icon: <CallToActionIcon />,
    configuration_id: "logger_webhook",
    configuration_label: "Logger Webhook",
  },
  {
    icon: <PendingActionsIcon />,
    configuration_id: "audit_webhook",
    configuration_label: "Audit Webhook",
  },
];

export const fieldsConfigurations: any = {
  region: [
    {
      name: "name",
      required: true,
      label: "Server Location",
      tooltip: 'Name of the location of the server e.g. "us-west-rack2"',
      type: "string",
      placeholder: "e.g. us-west-rack-2",
    },
    {
      name: "comment",
      required: false,
      label: "Comment",
      tooltip: "You can add a comment to this setting",
      type: "comment",
      placeholder: "Enter custom notes if any",
    },
  ],
  compression: [
    {
      name: "extensions",
      required: false,
      label: "Extensions",
      tooltip:
        'Extensions to compress e.g. ".txt",".log" or ".csv", you can write one per field',
      type: "csv",
      placeholder: "Enter an Extension",
      withBorder: true,
    },
    {
      name: "mime_types",
      required: false,
      label: "Mime Types",
      tooltip:
        'Mime types e.g. "text/*","application/json" or "application/xml", you can write one per field',
      type: "csv",
      placeholder: "Enter a Mime Type",
      withBorder: true,
    },
  ],
  api: [
    {
      name: "requests_max",
      required: false,
      label: "Requests Max",
      tooltip: "Maximum number of concurrent requests, e.g. '1600'",
      type: "number",
      placeholder: "Enter Requests Max",
    },
    {
      name: "cors_allow_origin",
      required: false,
      label: "Cors Allow Origin",
      tooltip: "List of origins allowed for CORS requests",
      type: "csv",
      placeholder: "Enter allowed origin e.g. https://example.com",
    },
    {
      name: "replication_workers",
      required: false,
      label: "Replication Workers",
      tooltip: "Number of replication workers, defaults to 100",
      type: "number",
      placeholder: "Enter Replication Workers",
    },
    {
      name: "replication_failed_workers",
      required: false,
      label: "Replication Failed Workers",
      tooltip:
        "Number of replication workers for recently failed replicas, defaults to 4",
      type: "number",
      placeholder: "Enter Replication Failed Workers",
    },
  ],
  heal: [
    {
      name: "bitrotscan",
      required: false,
      label: "Bitrot Scan",
      tooltip:
        "Perform bitrot scan on disks when checking objects during scanner",
      type: "on|off",
    },
    {
      name: "max_sleep",
      required: false,
      label: "Max Sleep",
      tooltip:
        "Maximum sleep duration between objects to slow down heal operation. eg. 2s",
      type: "duration",
      placeholder: "Enter Max Sleep duration",
    },
    {
      name: "max_io",
      required: false,
      label: "Max IO",
      tooltip:
        "Maximum IO requests allowed between objects to slow down heal operation. eg. 3",
      type: "number",
      placeholder: "Enter Max IO",
    },
  ],
  scanner: [
    {
      name: "delay",
      required: false,
      label: "Delay multiplier",
      tooltip: "Scanner delay multiplier, defaults to '10.0'",
      type: "number",
      placeholder: "Enter Delay",
    },
    {
      name: "max_wait",
      required: false,
      label: "Max Wait",
      tooltip: "Maximum wait time between operations, defaults to '15s'",
      type: "duration",
      placeholder: "Enter Max Wait",
    },
    {
      name: "cycle",
      required: false,
      label: "Cycle",
      tooltip: "Time duration between scanner cycles, defaults to '1m'",
      type: "duration",
      placeholder: "Enter Cycle",
    },
  ],
  etcd: [
    {
      name: "endpoints",
      required: true,
      label: "Endpoints",
      tooltip:
        'List of etcd endpoints e.g. "http://localhost:2379", you can write one per field',
      type: "csv",
      placeholder: "Enter Endpoint",
    },
    {
      name: "path_prefix",
      required: false,
      label: "Path Prefix",
      tooltip: 'Namespace prefix to isolate tenants e.g. "customer1/"',
      type: "string",
      placeholder: "Enter Path Prefix",
    },
    {
      name: "coredns_path",
      required: false,
      label: "Coredns Path",
      tooltip: 'Shared bucket DNS records, default is "/skydns"',
      type: "string",
      placeholder: "Enter Coredns Path",
    },
    {
      name: "client_cert",
      required: false,
      label: "Client Cert",
      tooltip: "Client cert for mTLS authentication",
      type: "string",
      placeholder: "Enter Client Cert",
    },
    {
      name: "client_cert_key",
      required: false,
      label: "Client Cert Key",
      tooltip: "Client cert key for mTLS authentication",
      type: "string",
      placeholder: "Enter Client Cert Key",
    },
    {
      name: "comment",
      required: false,
      label: "Comment",
      tooltip: "You can add a comment to this setting",
      type: "comment",
      multiline: true,
      placeholder: "Enter custom notes if any",
    },
  ],
  logger_webhook: [
    {
      name: "endpoint",
      required: true,
      label: "Endpoint",
      type: "string",
      placeholder: "Enter Endpoint",
    },
    {
      name: "auth_token",
      required: true,
      label: "Auth Token",
      type: "string",
      placeholder: "Enter Auth Token",
    },
  ],
  audit_webhook: [
    {
      name: "endpoint",
      required: true,
      label: "Endpoint",
      type: "string",
      placeholder: "Enter Endpoint",
    },
    {
      name: "auth_token",
      required: true,
      label: "Auth Token",
      type: "string",
      placeholder: "Enter Auth Token",
    },
  ],
};

export const removeEmptyFields = (formFields: IElementValue[]) => {
  const nonEmptyFields = formFields.filter((field) => field.value !== "");

  return nonEmptyFields;
};

export const selectSAs = (
  e: React.ChangeEvent<HTMLInputElement>,
  setSelectedSAs: Function,
  selectedSAs: string[]
) => {
  const targetD = e.target;
  const value = targetD.value;
  const checked = targetD.checked;

  let elements: string[] = [...selectedSAs]; // We clone the selectedSAs array
  if (checked) {
    // If the user has checked this field we need to push this to selectedSAs
    elements.push(value);
  } else {
    // User has unchecked this field, we need to remove it from the list
    elements = elements.filter((element) => element !== value);
  }
  setSelectedSAs(elements);
  return elements;
};
