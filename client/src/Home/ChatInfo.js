import * as React from "react";
import { useEffect, useState } from "react";
import { getChatRoomInfo } from "../api/api";
import { useNavigate } from "react-router-dom";
import Card from "@mui/material/Card";
import CardActions from "@mui/material/CardActions";
import CardContent from "@mui/material/CardContent";
import CardMedia from "@mui/material/CardMedia";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";
import { Alert } from "@mui/material";
import earth from "../static/images/earth.jpg";

export default function ChatInfo() {
  const navigate = useNavigate();
  const [showAlert, setShowAlert] = useState({
    status: "",
    value: "",
  });
  const [chatRoomInfo, setChatRoomInfo] = useState({ online: 0 });
  useEffect(() => {
    const fetchChatInfo = async () => {
      const token = localStorage.getItem("token");
      try {
        const res = await getChatRoomInfo(token);
        setChatRoomInfo({ online: res.data.online });
      } catch (err) {
        let errResponse = err.message;
        if (err.hasOwnProperty("response")) {
          errResponse = err.response.data.error;
        }
        setShowAlert({
          status: "error",
          value: `Failed to display chat room info: ${errResponse}`,
        });
      }
    };

    fetchChatInfo().catch(console.error);
  }, []);

  const handleJoin = () => {
    navigate("/chatroom");
  };

  return (
    <Card sx={{ width: "25%" }}>
      <CardMedia sx={{ height: 300 }} image={earth} title="chat room" />
      <CardContent>
        <Typography gutterBottom variant="h5" component="div">
          Chat Room
        </Typography>
        <Typography variant="body2" color="text.secondary">
          Online: {chatRoomInfo.online}
        </Typography>
      </CardContent>
      <CardActions>
        <Button size="large" onClick={handleJoin}>
          Join
        </Button>
      </CardActions>
      {showAlert.status && (
        <Alert severity={showAlert.status}>{showAlert.value}</Alert>
      )}
    </Card>
  );
}
