import userSchema from "../schema/userSchema";

import mongoose from "mongoose";

// creating and exporting the user model
export default mongoose.model(`user`, userSchema);
