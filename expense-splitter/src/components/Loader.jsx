import React from "react";
import styles from "./Loader.module.css";

export function Loader() {
  return (
    <div className={styles.loaderContainer}>
      <div className={styles.loader}></div>
    </div>
  );
}
