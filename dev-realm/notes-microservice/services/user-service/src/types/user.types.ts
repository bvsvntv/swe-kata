type CreateProfileType = {
    firstName: string;
    lastName: string;
    bio?: string;
    avatarURL?: string;
};

type UpdateProfileType = Partial<CreateProfileType>;

export { CreateProfileType, UpdateProfileType };
