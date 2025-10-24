import styles from "./style.module.scss";

export const RegisterPage = () => {
    return (
        <div className={styles.register_page}>
            <div className={styles.register_container}>
                <div className={styles.register_header}>
                    <h1>Create an Account</h1>
                    <p>Create an account and start managing your databases throughout the world.</p>
                </div>

                <form className={styles.register_form}>
                    <div className={styles.form_row}>
                        <div className={styles.form_group}>
                            <label htmlFor="firstName">First name</label>
                            <input
                                type="text"
                                id="firstName"
                                placeholder="Max"
                            />
                        </div>

                        <div className={styles.form_group}>
                            <label htmlFor="lastName">Last name</label>
                            <input
                                type="text"
                                id="lastName"
                                placeholder="Musterman"
                            />
                        </div>
                    </div>

                    <div className={styles.form_row}>
                        <div className={styles.form_group}>
                            <label htmlFor="email">Email Address</label>
                            <input
                                type="email"
                                id="email"
                                placeholder="mmusterman@mail.com"
                            />
                        </div>
                    </div>

                    <div className={styles.form_row}>
                        <div className={styles.form_group}>
                            <label htmlFor="password">Password</label>
                            <input
                                type="password"
                                id="password"
                                placeholder="••••••••••"
                            />
                        </div>

                        <div className={styles.form_group}>
                            <label htmlFor="confirmPassword">Confirm Password</label>
                            <input
                                type="password"
                                id="confirmPassword"
                                placeholder="••••••••••"
                            />
                        </div>
                    </div>

                    <div className={styles.checkbox_group}>
                        <input
                            type="checkbox"
                            id="terms"
                        />
                        <label htmlFor="terms">I agree the terms and conditions</label>
                    </div>

                    <div className={styles.form_actions}>
                        <button type="submit" className={styles.btn_primary}>
                            Sign Up
                        </button>
                        <button type="button" className={styles.btn_secondary}>
                            Sign In
                        </button>
                    </div>
                </form>
            </div>
        </div>
    );
};